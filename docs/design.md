# Basic Designs

## Single User sceninario

### Data Structures

#### structs

- `todo`
  - id: number
  - title: string, fixed-length
  - status: Status
  - hidden: boolean
  - parents: map of `edge`
  - children: map of `edge`
  - dayRoots: map of `edge`
  - milestons: map of `Milestone`
- `treeRoot` inherited from `todo`
  - color: #xxxxxx
  - order: int32
- `dayRoot` inherited from `todo`
  - date: Time
- `Milestone` inherited from `todo`
  - title: string
  - start: Time
  - end: Time
- `edge`
  - id: number
  - parent: reference for `todo`
  - son: reference for `todo`
  - root: reference for `rootTodo`
  - order: int32

#### category of roots

- TagRoot
  - dayRoot, Milestone, Draft, Selection
  - map of `edge`
- TreeRoot
