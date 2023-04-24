<script lang="ts">
    import Drawer, {
        AppContent,
        Content,
    } from "@smui/drawer";
    import List, { Item, Text, Separator, Meta } from "@smui/list";
    import IconButton from "@smui/icon-button";
    import NewNote from "./NewNote.svelte";
    import NotesPage from "./NotesPage.svelte";
	import { afterUpdate, createEventDispatcher } from 'svelte';
    import type { backend } from "../../wailsjs/go/models";
    import NoteSearch from "./NoteSearch.svelte";

    export let active: backend.Note;
    export let notes: backend.Note[];

    let displayNotes = notes;
    console.log(notes);
    console.log(displayNotes);

    let open = true;
    let search = "";
    let dispatch = createEventDispatcher();

    function setActive(value: backend.Note) {
        dispatch ("activate", {
            page: value
        });
    }

    function newNote() {
        dispatch ("new");
    }

    function resetNotes() {
        displayNotes = notes;
    }

    afterUpdate(() => {
        if (search.length === 0) {
            resetNotes();
        }
    });

    function onSearch(event: CustomEvent<any>) {
        search = event.detail.search;
        if (search.length > 0) {
            console.log("searching...")
            displayNotes = notes.filter((note) => {
                return (
                    note.title
                        .toLowerCase()
                        .includes(event.detail.search.toLowerCase()) ||
                    note.summary
                        .toLowerCase()
                        .includes(event.detail.search.toLowerCase()) ||
                    note.tags
                        .map((tag) => tag.name)
                        .join(", ")
                        .toLowerCase()
                        .includes(event.detail.search.toLowerCase())
                );
            });
            console.log(notes);
        } else {
            console.log("reseting...")
            resetNotes();
        }
    }

    function onSaved() {
        dispatch ("saved");
        setTimeout(() => {
            resetNotes();
        }, 200);
    }

    function onUpdate(event: CustomEvent<any>) {
        dispatch ("update", {
            note: event.detail.note
        });
        setTimeout(() => {
            resetNotes();
        }, 200);
    }

    function onDelete(id: number) {
        dispatch ("delete", {
            id: id
        });
        setTimeout(() => {
            resetNotes();
        }, 200);
    }

</script>

<div class="drawer-container">
    <Drawer variant="dismissible" bind:open>
        <Content>
            <List>
                <Item
                    href="javascript:void(0)"
                    on:click={() => newNote()}
                    activated={active.id === -1}
                >
                    <Text>New Note</Text>
                </Item>
                <Separator />
                <NoteSearch on:search={(event) => onSearch(event)} />
                {#each displayNotes as item}
                    <Item
                        href="javascript:void(0)"
                        on:click={() => setActive(item)}
                        activated={active === item}
                    >
                        <Text>{item.title}</Text>
                        <Meta>
                            <IconButton class="material-icons" on:click={() => onDelete(item.id)}
                                >delete</IconButton
                              >
                            </Meta>
                    </Item>
                {/each}
            </List>
        </Content>
    </Drawer>

    <AppContent class="app-content">
        <main class="main-content">
            {#if active.title === "New Note"}
                <NewNote on:saved={() => onSaved()} />
            {:else}
                <NotesPage bind:note={active} on:update={(event) => onUpdate(event)} />
            {/if}
        </main>
    </AppContent>
</div>

<style>

</style>
