<script lang="ts">
    import { Separator } from "@smui/list";
    import { formatContent } from "../utils/Files";
    import { createEventDispatcher } from "svelte";
    import Textfield, { Input, Textarea } from "@smui/textfield";
    import HelperText from "@smui/textfield/helper-text";
    import IconButton from "@smui/icon-button";
    import Paper, { Title, Subtitle, Content } from "@smui/paper";
    import type NotchedOutline from "@smui/notched-outline";
    import FloatingLabel from "@smui/floating-label";
    import LineRipple from "@smui/line-ripple";
    import FormField from "@smui/form-field";
    import Switch from "@smui/switch";  
    import { backend } from "../../wailsjs/go/models";
    import Welcome from "../pages/Welcome.svx";

    export let note: backend.Note;

    let tagsValue = note.tags? note.tags.map((tag) => tag.name).join(", ") : "";

    let tagsInput: Input;
    let tagsFloatingLabel: FloatingLabel;
    let tagsLineRipple: LineRipple;

    let tags = [] as backend.Tag[];
    
    $: {
        if (tagsValue.length > 0) {
            tags = tagsValue.split(",").map((tag) => {
                return new backend.Tag({name: tag.trim()});
            });
        }
    }
    let summaryInput: Textarea;
    let summaryNotchedOutline: NotchedOutline;

    let contentInput: Textarea;
    let contentNotchedOutline: NotchedOutline;

    let content: string;
    let dispatch = createEventDispatcher();
    let editMode = false;
    let button = "edit";

    $: button = editMode? "arrow_back" : "edit";

    $: content = formatContent(note.content);

    function updateNote() {
        dispatch("update", {
            note: note,
        });
    }

</script>

<div class="paper-container">
    {#if note.id === -1}
        <Welcome />
    {:else}
        <IconButton class="material-icons" on:click={() => editMode = !editMode}>{button}</IconButton>
        {#if editMode}
        <IconButton class="material-icons" on:click={() => updateNote()}>save</IconButton>
            <Paper color="primary" variant="outlined">
                <Title>
                    <Textfield
                        bind:value={note.title}
                        placeholder="Title"
                        style="width: 98%; margin: 0.5rem;"
                        />
                </Title>
                <Textfield
                bind:input={tagsInput}
                bind:floatingLabel={tagsFloatingLabel}
                bind:lineRipple={tagsLineRipple}
                style="width: 98%; margin: 0.5rem;"
            >
                <FloatingLabel
                    bind:this={tagsFloatingLabel}
                    for="input-manual-a"
                    slot="label">Tags</FloatingLabel
                >
                <Input
                    bind:this={tagsInput}
                    bind:value={tagsValue}
                    id="input-manual-a"
                />
                <HelperText>Separate tags with a comma</HelperText>
                <LineRipple bind:this={tagsLineRipple} slot="ripple" />
            </Textfield>
        <pre>
            <div class="tag-container">
                {#each tags as tag}
                <div class="tag">
                    {tag.name}
                </div>
                {/each}
            </div>
        </pre>
        <Subtitle>
                    <div>
                    <Textfield
                        bind:input={summaryInput}
                        bind:NotchedOutline={summaryNotchedOutline}
                        placeholder="Summary"
                        style="width: 98%; height: 100px; margin: 0.5rem;"
                        textarea
                    >
                        <Textarea
                            bind:this={summaryInput}
                            bind:value={note.summary}
                        />
                    </Textfield>
                    </div>
                </Subtitle>
                <FormField align="end">
                    <Switch bind:checked={note.important} id="input-switch-important" />
                    <span slot="label">Important</span>
                </FormField>
                <Separator />
                <br />
                <Content>
                    <Textfield
                        bind:input={contentInput}
                        bind:NotchedOutline={contentNotchedOutline}
                        placeholder="Content"
                        style="text-align: top; width: 98%; height: 300px; margin: 0.5rem;"
                        textarea
                    >
                        <Textarea
                            bind:this={contentInput}
                            bind:value={note.content}
                        />
                    </Textfield>
                </Content>
            </Paper>
        {:else}
        <Paper color="primary" variant="outlined">
            <h1>{note.title}<IconButton  class="material-icons" on:click={() => editMode = !editMode}>pencil</IconButton></h1>
            <pre>
                <div class="tag-container">
                    {#each note.tags as tag}
                    <div class="tag">
                        {tag.name}
                    </div>
                    {/each}
                </div>
            </pre>
            <Subtitle>
                created: {note.created_date}
                <br />
                updated: {note.updated_date}
            </Subtitle>
            <Separator />
            <br />
            <Content>{note.summary}</Content>
            <br />
            <Separator />
            <br />
            <Content>{@html content}</Content>
        </Paper>
        {/if}
    {/if}

</div>

<style>
    .paper-container {
        padding: 1rem;
        height: 100%;
        width: 90%;
    }
    .tag-container {
        display: flex;
        max-width: 98%;
        flex-wrap: wrap;
        }
    .tag {
        display: flex;
        background-color: #222;
        border-radius: 100px;
        margin: 2px;
        font-size: medium;
        overflow-wrap: scroll;
        height: 20px;
        width: 60px;
        max-width: 75px;
        justify-content: center;
        text-align: center;
    }
</style>