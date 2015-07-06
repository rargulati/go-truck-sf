Experience: Two months experience casually writing Go (automating system tasks). Some code review.
Problem: As a person in San Francisco, I'd love to know where the closest, open food trucks are. 
I'd also enjoy the ability to sort by a variety of parameters (type)
Solution: Api first backend that allows the user to hit the following endpoints to accomplish the above tasks.
Many different clients can access the ability to view all food trucks in the default radius and from there query based off
what makes sense as default or for user.

As I haven't provided a rich-front end experience from the user to use the app with,
I focused on an api-first design where a developer more capable (or future me with more time)
in front-end development can consume data by querying 4 simple endpoints.

Trade-offs:
As this was quick and dirty, as of now, no rich front-end experience. Simple Go templating.
A look at the socialpoll example (influenced by an exercise in a golang book) shows what a golang project
that is built for scale would look like.

Missing:
Tests! In the middle of figureing out Golang testing conventions, but will be testing the following cases:
1) A valid request to the server results in a well-formed response (200) or an http error (404)
2) If the server responds with 200, it returns a slice of structs of type Facility.
3) That queries about the data set the correct url params in the request from server to server
4) The above queries also return well-formed ata (slice of structus of type Facility)

Link to other code I'm proud of: some haskell stuff, but I really need to open source more.
Link to resume:
Link to app:
