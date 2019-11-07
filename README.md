# http-fake
Starts fake configurable http server for development

Simple http server which can be configured through file

### Config parameters:
- port: server port
- corsEnabled: enable servers CORS
- routes: map[string][Route](#Route) - map of routes 

### Route:
- method: http method
- content: [Content](#Content)
- headers: map[string]string

### Content:
- type: string
- body: string
- link: string
- random: [][RandomItem](#RandomItem)

### RandomItem:
- body: string

Config example:
```
port: 3333
corsEnabled: true
routes:
  /:
    # route accept only http get method
    method: get 
    content: 
      # response will be get from body property
      type: inline
      body: '{"data": "hello world"}'
  /link: 
    method: post
    content: 
      # response will be get from link property
      # server will search file named './index.json' 
      # and it's content will be send as response body
      type: link
      link: ./index.json 
  /random: 
    method: put
    content: 
      # response body will be get as random value from 'random' array
      type: random
      random:
        - body: '{"data": "hello world 1"}'
        - body: '{"data": "hello world 2"}'
```