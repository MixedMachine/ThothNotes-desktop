# Thoth Notes Desktop app

---

## About
The goal of this app is to provide a simple way to take notes and organize them. 
It is built with Wails 2 and Svelte. The notes are saved in a local database,
and the app is designed to be used offline. \

## Future Features
- [X] Add a search bar
- [X] Add tags / topics
- [X] Separate created and updated dates
    - [ ] Order by created date
    - [ ] Order important notes feed by updated date
- [ ] Add image list
- [ ] Add related notes list
- [ ] Export notes to markdown file or pdf

## Bugs / Refactoring
- [ ] Getting the dates to show up right
- [ ] When page is up, deleting from list leaves page up until switching pages
- [ ] clean up tags display into its own component
- [ ] Tag elements adjust to text size
- [ ] created at is broken
- [ ] Break up the note page into components
- [ ] Create form component for new and update to reduce code duplication
- [ ] Refactor the database or store to not duplicate tags
- [ ] Update notes better

## Development
### Requirements
- [Wails 2](https://wails.app/gettingstarted/)
- [Node.js](https://nodejs.org/en/download/)
- [Go](https://golang.org/doc/install)

### Setup
1. Clone the repo
2. Run `make run` to start the app in development mode
- exe will be in the `build/bin` folder

### Building installer for windows
1. Run `make prod` to build the app installer
- installer will be in the `build/bin` folder
- pre-built installer can also be found [here](https://thoth-notes.s3.us-west-2.amazonaws.com/ThothNotes-amd64-installer.exe)

