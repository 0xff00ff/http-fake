# http-fake
Starts fake configurable http server for development

Simple http server which can be configured through file

### Config parameters:
- port:        string #server port
- corsEnabled: bool #enable servers CORS
- routes:      map[string][Route](#Route) - map of routes 

### Route:
- method:  string #http method
- content: [Content](#Content)

### Content:
- type:     string # inline | file
- algorithm string # static | random | condition
- body:     any    # string | [ConditionalBody](#ConditionalBody)

#### ConditionalBody:
- query  string
- value  string
- body   string
- source string

### RandomItem:
- body: string

Config example:
```
port: 3333
corsEnabled: true
routes:

  /inline:
    method: get
    content: 
      type: inline
      body: '{"data": "hello world"}'

  /link:
    method: get
    content: 
      type: file 
      body: ./asset.json

  /random:
    method: get
    content: 
      type: inline
      algorithm: random
      body:
        - '{"data": "hello world 1"}'
        - '{"data": "hello world 2"}'
        - '{"data": "hello world 3"}'
        - '{"data": "hello world 4"}'
      
  /random/:qwe:
    method: get
    content: 
      type: inline
      algorithm: random
      body:
        - '{"data": "hello world 1"}'
        - '{"data": "hello world 2"}'
        - '{"data": "hello world 3"}'
        - '{"data": "hello world 4"}'

  /random-file:
    method: get
    content: 
      type: file
      algorithm: random
      body: 
        - './dog.json'
        - './cat.json'
        - './fish.json'
  
  /condition:
    method: post
    content:
      type: inline
      algorithm: condition
      body:
        - source: body
          query: 'a[0]'
          value: "cat"
          body: '{"type": "cat"}'
        - source: body
          query: 'a[0]'
          value: "dog"
          body: '{"type": "dog"}'

  /condition/animal/:animal-type:
    method: get
    content:
      type: file
      algorithm: condition
      body:
        - source: path
          query: 'animal-type'
          value: "cat"
          body: 'cat.json'
        - source: path
          query: 'animal-type'
          value: "dog"
          body: 'dog.json'
```