
## What is Yaml ?

YAML (a recursive acronym for “YAML Ain’t Markup Language”) is a data **serialization** language designed to be   
human-friendly and work well with modern programming languages for common everyday tasks.  
YAML is a superset of JSON, which means that any valid JSON file is also a valid YAML file.  

### Yaml data structure  

Two commonly used data structures in yaml are : map & list  

YAML Map ::  
Map is used to store the **name** (or key) **value** pair.  

Ex- we have two values, v1 and Pod, mapped to two keys, apiVersion, and kind.  

```
apiVersion: v1
kind: Pod
```

The JSON equivalent of above,  

```
{
   "apiVersion": "v1",
   "kind": "Pod"
}
```

Example of more complicated structures by creating a key that _maps to another map_, rather than a string, as in:

```
apiVersion: v1
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
```

Note that, the YAML processor knows how all of these pieces relate to each other because we’ve indented the lines.  
In this example we’ve used 2 spaces for readability, but the number of spaces doesn’t matter — as long as it’s at least 1,  
 and as long as you’re CONSISTENT. For example, name and labels are at the same indentation level, so the processor knows  
 they’re both part of the same map; it knows that app is a value for labels because it’s indented further.  
 
So if we were to translate this to JSON, it would look like this:

```
{
  "apiVersion": "v1",
  "kind": "Pod",
  "metadata": {
               "name": "rss-site",
               "labels": {
                          "app": "web"
                         }
              }
}
```

YAML List ::  

YAML lists are literally a sequence of objects.  For example:  

```
args:
  - sleep
  - "1000"
  - message
  - "Bring back Firefly!"
```

As you can see here, we can have virtually any number of items in a list, which is defined as items that start with a dash (-) indented from the parent.  
So in JSON, this would be:  

```
{
   "args": ["sleep", "1000", "message", "Bring back Firefly!"]
}

```

And of course, members of the list can also be maps:

```
apiVersion: v1
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
spec:
  containers:
    - name: front-end
      image: nginx
      ports:
        - containerPort: 80
    - name: rss-reader
      image: nickchase/rss-php-nginx:v1
      ports:
        - containerPort: 88
```

So as we can see here, we have a list of container “objects”, each of which consists of a name, an image,  
and a list of ports (It might also include network information). Each list item under ports is itself a map  
that lists the containerPort and its value.  

For completeness, let’s quickly look at the JSON equivalent:  

```
{
  "apiVersion": "v1",
  "kind": "Pod",
  "metadata": {
                "name": "rss-site",
                "labels": {
                            "app": "web" 
                          }
              },
   "spec": {
              "containers": [ 
                              { 
                                "name": "front-end",
                                "image": "nginx",
                                "ports": [
                                           {
                                              "containerPort": "80"
                                           }
                                         ]
                              }, 
                              {
                                "name": "rss-reader",
                                "image": "nickchase/rss-php-nginx:v1",
                                "ports": [
                                           {
                                              "containerPort": "88"
                                           }
                                         ]
                              }
                            ]
           }
}
```

So let’s review.  We have:  

- maps, which are groups of name-value pairs
- lists, which are individual items
- maps of maps
- maps of lists
- lists of lists
- lists of maps  

Basically, whatever structure you want to put together, you can do it with those two structures !!

More example of creating a pod in kubernetes using a yaml file:  

```
apiVersion: v1
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
spec:
  containers:
    - name: front-end
      image: nginx
      ports:
        - containerPort: 80
    - name: rss-reader
      image: nickchase/rss-php-nginx:v1
      ports:
        - containerPort: 88
```



## Reference ::
[1]. [Yaml specification](https://yaml.org/)
