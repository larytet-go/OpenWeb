# OpenWeb Backend Programming Challenge

OpenWeb creates a helathy social networking layer over some of the worlds largest publishers, such as Fox News, TechCrunch, Yahoo, and more. 
We are on a mission to save online conversation - and our main product is our Conversation Module, which allows people to write messages about articles where the conversation is found.

## Instructions

In conversation, we need a way to represent our Message Stack that you see on the page. We have comments, replies, and replies to replies. Your mission, should you choose to accept it, is to build functionality that will allow us to store a tree of messages, with replies being connected to parents, etc.

We have started you off with a basic build:
- In `node.go`, you'll find a `node` struct, containing the actual message inside
- In `tree.go`, you'll find a `tree` struct, containing the slice of nodes, which should represent the Message Stack
- In `tree_test.go`, you'll find a few different tests which we will run to ensure that the data structure is correct
  - `TestNewTree` is the basic test, which will ensure a small data structure
  - `TestNewTreeLarge` contains a much bigger sample with more depth

## Bonus

Our product has billions of requests per day, and our scale means we must work in highly concurrent environments. Your bonus mission is to ensure that the data structure that you've created will fit the needs of a highly concurrent environment as well.

- In `concurrent_tree_test.go`, you'll find a few different tests which we will run to ensure that the data structure is correct
  - `TestNewTreeConcurrent` is the basic test, which will ensure a small data structure
  - `TestNewTreeLargeConcurrent` contains a much bigger sample with more depth


## Bonus #2

Now is the time to make that dream come true.

Turn your new functionality into an API service with 2 endpoints:
- `Add Comment` - adding new comment / reply (message) to a conversation_id.
- `Get Conversation` - get messages tree by conversation_id.

Don't forget! our biggest challenge is scale, so please make sure we can easily run and add new independent instances of your service.
Notes:
- You can use any DB / Data Store you know and like, there are a lot of already made dockers that run whatever DB you need.
- You are advised to use the “gin” framework for Golang. https://github.com/gin-gonic/gin
- Please add samples for API requests.


