<script>
    import UserSearch from './UserSearch.svelte';
    import {formStage, userData} from '$lib/stores';

    let currentStage = 1;

    let inputName = "";
    let error = "";

    let id = "";
    let names = []
    let guests = [];

    $: formStage.subscribe((stage) => {
        currentStage = stage;
    });

    $: userData.subscribe((data) => {
        id = data.id || "";
        names = data.names || [];
        guests = data.guests || [];
    });

</script>


<div>
    <div class="headerImage">
        <picture >
            <img src="/img/site-border-2.png" alt=""/>
        </picture>
    </div>
    <div style="justify-content: center; text-align: center; width:100%">
        <UserSearch/>
    </div>
</div>
