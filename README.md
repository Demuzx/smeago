# smeago
This is fork from github.com/BertGit/go-crawler repository. A Golang tool to generate sitemap xml for Server Side Rendering apps.
Now you can add user agent for crawler by -a flag, like: -a googlebot, etc.

![all text](http://orig14.deviantart.net/9c7b/f/2012/268/9/d/doodle__cute__gollum_by_agathexu-d5fu2mf.jpg)

## Install

```
go get github.com/Demuzx/smeago
```

## Example usage

```
smeago -o "www/" -h "http://example.com" -a googlebot
```

### Params

```
-h the host name to crawl
   default: http://localhost
-p the host port to crawl
   default: 80
-loc the host to be prefixed with the paths in the sitemap
   default: http://localhost
   example: -loc http://example.com
   <url>
    <loc>
      http://example.com/foo/bar
    </loc>
   </url>
-o the relative output directory for the sitemap.xml file
   default: <current directory>
-a set user-agent
   default: ""
   example: -a googlebot
```
