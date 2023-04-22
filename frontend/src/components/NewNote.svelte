<script lang="ts">
    import { addNote } from "../utils/Notes";
    import { createEventDispatcher } from "svelte";
    import Textfield, { Input, Textarea } from "@smui/textfield";
    import HelperText from "@smui/textfield/helper-text";
    import IconButton, { Icon } from "@smui/icon-button";
    import { mdiContentSaveOutline, mdiCloseThick } from "@mdi/js";
    import { Svg } from "@smui/common";
    import FormField from "@smui/form-field";
    import Switch from "@smui/switch";
    import FloatingLabel from "@smui/floating-label";
    import LineRipple from "@smui/line-ripple";
    import NotchedOutline from "@smui/notched-outline";
    import Chip, { Set, Text } from '@smui/chips';
    import {backend} from "../../wailsjs/go/models";

    let dispatch = createEventDispatcher();

    let titleValue = "";
    let titleInput: Input;
    let titleFloatingLabel: FloatingLabel;
    let titleLineRipple: LineRipple;

    let summaryValue = "";
    let summaryInput: Textarea;
    let summaryNotchedOutline: NotchedOutline;
    let summaryFloatingLabel: FloatingLabel;

    let importantValue = false;

    let contentValue = "";
    let contentInput: Textarea;
    let contentNotchedOutline: NotchedOutline;
    let contentFloatingLabel: FloatingLabel;

    let tagsValue = "";
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

    function reset() {
        titleValue = "";
        tagsValue = "";
        summaryValue = "";
        importantValue = false;
        contentValue = "";
        tags = [];

        titleInput.focus();
    }

    async function save() {
        console.log("Saving note...");

        await addNote(new backend.Note({
            title: titleValue,
            summary: summaryValue,
            important: importantValue,
            content: contentValue,
            tags: tags,
        }));

        setTimeout(function () {
            dispatch("saved");
        }, 100);

        reset();
    }
</script>

<div class="columns margins">
    <div>
        <Textfield
            bind:input={titleInput}
            bind:floatingLabel={titleFloatingLabel}
            bind:lineRipple={titleLineRipple}
            style="width: 98%; margin: 0.5rem;"
        >
            <FloatingLabel
                bind:this={titleFloatingLabel}
                for="input-manual-a"
                slot="label">Title</FloatingLabel
            >
            <Input
                bind:this={titleInput}
                bind:value={titleValue}
                id="input-manual-a"
            />
            <LineRipple bind:this={titleLineRipple} slot="ripple" />
        </Textfield>
    </div>
    <br />
    <div>
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
    </div>
    <pre>
        <div class="tag-container">
            {#each tags as tag}
            <div class="tag">
                {tag.name}
            </div>
            {/each}
        </div>
    </pre>
    <br />
    <div>
        <Textfield
            bind:input={summaryInput}
            bind:notchedOutline={summaryNotchedOutline}
            bind:floatingLabel={summaryFloatingLabel}
            style="width: 98%; height: 100px; margin: 0.5rem;"
            textarea
        >
            <NotchedOutline bind:this={summaryNotchedOutline} slot="label">
                <FloatingLabel
                    bind:this={summaryFloatingLabel}
                    for="input-manual-b">Summary</FloatingLabel
                >
            </NotchedOutline>
            <Textarea
                bind:this={summaryInput}
                bind:value={summaryValue}
                id="input-manual-b"
            />
        </Textfield>
    </div>
    <br />
    <div style="margin-left: 1em;">
        <FormField align="end">
            <Switch bind:checked={importantValue} id="input-manual-c" />
            <span slot="label">Important</span>
        </FormField>
    </div>
    <br />
    <div class="margins">
        <Textfield
            bind:input={contentInput}
            bind:notchedOutline={contentNotchedOutline}
            bind:floatingLabel={contentFloatingLabel}
            textarea
            style="width: 98%; height: 300px; margin: 0.5rem;"
        >
            <NotchedOutline bind:this={contentNotchedOutline} slot="label">
                <FloatingLabel
                    bind:this={contentFloatingLabel}
                    for="input-manual-d">Content</FloatingLabel
                >
            </NotchedOutline>
            <Textarea
                bind:this={contentInput}
                bind:value={contentValue}
                id="input-manual-d"
            />
        </Textfield>
    </div>
    <br />
    <div class="margins" style="align-items: center;">
        <IconButton class="material-icons" on:click={reset} size="button">
            <Icon component={Svg} viewBox="0 0 24 24">
                <path fill="currentColor" d={mdiCloseThick} />
            </Icon>
        </IconButton>
        <IconButton class="material-icons" on:click={save} size="button">
            <Icon component={Svg} viewBox="0 0 24 24">
                <path fill="currentColor" d={mdiContentSaveOutline} />
            </Icon>
        </IconButton>
    </div>
</div>

<style>
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
