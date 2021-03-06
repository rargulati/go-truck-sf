Experience:
------------
Two months experience casually writing Go (automating system tasks, very small scripts).

Problem
------------

As a person in San Francisco, I'd love to know where the closest, open food trucks are.
I'd also enjoy the ability to sort by a variety of parameters (type)

Solution:
------------
*Construction Zone*

The application has a main function that decodes the json representation of the data that we require.

Future
------------

(Eventually an) api-first backend that allows the user to hit a few endpoints to accomplish the above tasks.
Many different clients can access the ability to view all food trucks in the default radius and from there query based off
what makes sense as default or for user.

Trade-offs
------------

As this was quick and dirty (for now), no rich front-end experience or complex folder structure splitting out all the functions for webscale.
Simple Go templating and a few functions in a main package.

At this point I'm just levering standard libraries, though one big level up would be wrapping access to the api in a goroutine for concurrent access.
I'd also want to add a per-user lock (mutex) on their ability to filter queries as I don't want to exhaust the google api limits.

Missing:
------------

Tests! In the middle of figuring out Golang testing conventions, but will be testing the following cases:

1. A valid request to the server results in a well-formed response (200) or an http error (404)
2. If the server responds with 200, it returns a slice of structs of type Facility.
3. That queries about the data set the correct url params in the request from server to server
4. The above queries also return well-formed ata (slice of structus of type Facility)

Misc
------------

Link to other code I'm proud of: Most of my work is private and owned by others. www.github.com/rargulati for some dabblings.
<br />
Link to resume: www.angel.co/rargulati && www.linkedin.com/in/rargulati
<br />
Link to app: forthcoming
