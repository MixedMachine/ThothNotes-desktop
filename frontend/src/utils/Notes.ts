import { CreateNote, DeleteNoteByID, GetNotes, GetNoteByID, UpdateNoteByID, CreateNoteWithTags } from '../../wailsjs/go/backend/NotesHandler'
import type { backend } from '../../wailsjs/go/models'

export async function getNotes () {
    console.log('getNotes')
    let allNotes = [] as backend.Note[]
    await GetNotes().then((notes: backend.Note[]) => {
        allNotes = notes
        console.log(allNotes)
    })
    return allNotes
}

export async function getNote (id: number) {
    console.log('getNote')
    let note: backend.Note
    await GetNoteByID(id).then((note: backend.Note) => note = note)
    return note
}

export async function deleteNote (id: number) {
    console.log('deleteNote')
    await DeleteNoteByID(id)
}

export function updateNote (note: backend.Note) {
    console.log('updateNote')
    UpdateNoteByID(note.id, note.title, note.summary, note.content, note.important, note.tags)
}

export function addNote (note: backend.Note) {
    console.log('addNote')
    CreateNoteWithTags(note.title, note.summary, note.content, note.important, note.tags)
}
