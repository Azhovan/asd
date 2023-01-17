### What is a context? 
A context is a collection of categorized information.
for simplicity, you can imagine it as a flat file divided into headlines and each headline has a bunch of subsections. 
each section is simple data structure that keeps track of a goal and it's progress(if set by developer). 

context can be exported in a different ways
- A a flat readme file 
- As a git repository. this is particular if you'd like to share progress of a project with your colleagues. 

### Context Spec
Contexts have a structure as followings: 
```golang 
type Context struct {
    Title string
    Descriptions []string
    Headlines []Headline
}

// Headline represents a single goal tha breaked down into actions items.
type Headline struct {
	Title       string
	Done        bool
	ActionItems []ActionItem
}

// ActionItem represents a single action that is part of a headline.
type ActionItem struct {
	Title string
	Done  bool
}
```


 
