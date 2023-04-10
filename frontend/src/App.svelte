<script lang="ts">
import {
    getNotes,
    updateNote,
    deleteNote
} from "./utils/Notes";
import {
    onMount
} from "svelte";
import type {
    backend
} from '../../wailsjs/go/models';
import NotesFeed from "./components/NotesFeed.svelte";
import NotesList from "./components/NotesList.svelte";
import LayoutGrid, {
    Cell
} from "@smui/layout-grid";

let defaultPage = {
    id: -1,
    title: "Welcome",
};

let newNotePage = {
    id: -1,
    title: "New Note",
}

let active = defaultPage;
let notes = [] as backend.Note[];
let importantNotes = [] as backend.Note[];

function GetNotes() {
    setTimeout(async () => {
        notes = await getNotes();
    }, 100);
}

function DeleteNote(id: number) {
    deleteNote(id);
    GetNotes();
}

function UpdateNote(note: backend.Note) {
    updateNote(note);
    GetNotes();
}

onMount(() => {
    GetNotes();
});

$: importantNotes = notes.filter((note) => note.important);
</script>

<main>
    <div class="notes-header">
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <h1 class="app-header-title" on:click={ () => active = defaultPage }>Thoth Notes</h1>
    </div>
    <div class="notes-container">
        <LayoutGrid class="notes-container">
            <Cell span={3}>
                <div class="notes-list">
                    <NotesList
                        bind:active
                        bind:notes
                        on:new={() => active = newNotePage}
                        on:activate={(event) => active=event.detail.page}
                        on:saved={() => GetNotes()}
                        on:update={(event) => UpdateNote(event.detail.note)}
                        on:delete={(event) => DeleteNote(event.detail.id)}
                        />
                        </div>
                        </Cell>
                        <Cell span={3}>
                            <div class="notes-feed">
                                <NotesFeed
                                    bind:notes={importantNotes}
                                    on:activate={(event) => active=event.detail.page}
                                    />
                                    </div>
                                    </Cell>
                                    </LayoutGrid>
                                    </div>
                                    </main>

<style>
.notes-header {
    position: fixed;
    top: 0;
    width: 100%;
    padding: 0.1rem;
}

.app-header-title {
    font-size: 3rem;
    font-weight: 500;
    margin: 0;
    display: inline-block;
    padding: 0.5rem;
    cursor: pointer;
}

.notes-container {
    position: fixed;
    top: 6rem;
    height: 100%;
}

.notes-list {
    height: 85vh;
    position: fixed;
    left: 0;
    width: 75vw;
    overflow-y: scroll;
}

.notes-feed {
    height: 85vh;
    position: fixed;
    right: 0;
    width: 25vw;
    overflow-y: scroll;
}

.notes-list::-webkit-scrollbar {
    width: 1em;
}

.notes-list::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
}

.notes-list::-webkit-scrollbar-thumb {
    background-color: darkgrey;
    outline: 1px solid slategrey;
}

.notes-feed::-webkit-scrollbar {
    width: 1em;
}

.notes-feed::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
}

.notes-feed::-webkit-scrollbar-thumb {
    background-color: darkgrey;
    outline: 1px solid slategrey;
}
</style>
