<script lang="ts">
    import List, {
        Item,
        Text,
        PrimaryText,
        SecondaryText,
    } from "@smui/list";
    import { getNotes } from "../utils/Notes";
    import { createEventDispatcher } from 'svelte';

    export let notes;

    let list: InstanceType<typeof List>;
    let dispatch = createEventDispatcher();

    function setActive(value: Note) {
        dispatch ("activate", {
            page: value
        });
    }

</script>

<div>
    <div class="head"><h3>Important Notes Feed</h3></div>
    <List
        class="notes-list"
        bind:this={list}
        twoLine
        avatarList
        singleSelection
    >
        {#each notes as item, i}
            <Item
                on:SMUI:action={() => (setActive(item))}
            >
                <Text>
                    <PrimaryText>{item.title}</PrimaryText>
                    <SecondaryText>{item.summary}</SecondaryText>
                </Text>
            </Item>
        {/each}
    </List>
</div>

<style>
    * :global(.demo-list) {
        max-width: 600px;
    }

    .head {
        margin: 0 0 1rem 0;
        text-align: center;
    }
</style>
