package smeago

import (
	"log"
	"net/http"
	"time"
)

type Job struct {
	ID         int
	Path       string
	Links      []string
	Completed  bool
	RetryCount int
}

// NewJob creates a new Job
func NewJob(id int, path string) *Job {
	return &Job{
		ID:   id,
		Path: path,
	}
}

type Crawler struct {
	Domain     string
	Results    chan Job
	Retries    chan Job
	Headers    map[string]string
	MaxRetries int
	ReqTimeout time.Duration
}

// NewCrawler creates a crawler for the given domain
func NewCrawler(d string, reqTimeout time.Duration, maxRetries int) *Crawler {
	c := &Crawler{}
	c.Domain = d
	c.Results = make(chan Job)
	c.Retries = make(chan Job)
	c.Headers = make(map[string]string)
	c.MaxRetries = maxRetries
	c.ReqTimeout = reqTimeout
	return c
}

func (c *Crawler) AddHeader(key, value string) {
	c.Headers[key] = value
}

// Crawl the job path and retries in case of failures
func (c *Crawler) Crawl(j Job) {
	var (
		err error
		r   *Result
	)
	link := c.Domain + j.Path

	if j.RetryCount > 0 {
		log.Printf("Retrying (%d): %s\n", j.RetryCount, link)
		if j.RetryCount > c.MaxRetries {
			j.Completed = true
			c.Results <- j
		}
	} else {
		log.Println("Visiting:", link)
	}

	client := http.Client{
		Timeout: c.ReqTimeout,
	}
	req, _ := http.NewRequest("GET", link, nil)
	for k, v := range c.Headers {
		log.Printf("add custom header: %v: %v", k, v)
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.Retries <- j
		return
	}
	defer resp.Body.Close()

	n := int(resp.ContentLength)
	if n > 0 {
		r, err = ReadStringSize(resp.Body)
	} else {
		r, err = ReadString(resp.Body)
	}

	if err != nil {
		log.Println(err)
		c.Results <- j
		return
	}

	j.Links = r.Links
	j.Completed = true
	c.Results <- j
}
