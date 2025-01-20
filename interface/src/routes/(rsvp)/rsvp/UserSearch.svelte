<script>
    import { formStage, userData } from '$lib/stores';
    import GuestInput from "./GuestInput.svelte";
    import DinnerSelect from "./DinnerSelect.svelte";

    let name = '';
    let error = '';
    let currentStage = 1;

    async function handleSubmit() {
        try {
            const response = await fetch('/rsvp/fetch_user', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name }),
            });

            if (response.ok) {
                const data = await response.json();
                userData.set(data); // Populate the user data store
                console.log($userData);
                currentStage++;
                formStage.set(2); // Move to the next stage
            } else {
                error = 'User not found. Please try again.';
            }
        } catch (err) {
            error = 'An error occurred. Please try again.';
        }
    }

    const handleNext = () => {
        //Todo: Do some sorta validation here
        currentStage++;
        console.log($userData.guests);
    }
    
    const handleBack = () => {
        currentStage--;
        //Todo: if stage == 1, clear userData;
    }

    async function handleSave() {

        console.log("In handleSave")
        console.log($userData)
        try {
            const response = await fetch('/rsvp/update_user', {
                method: "POST",
                headers: { "Content-Type": "application/json"},
                body: JSON.stringify($userData)
            });

            if (response.ok) {
                currentStage++
            }
        } catch (err) {
            error = 'An error occurred. Please try again.';
        }
        // Todo: Get the actual saving done here
        console.log($userData.guests);
    }
</script>


<div class="rsvpForm">
    {#if currentStage === 1}
        <h2>Marina & Edward</h2>
        <div class="rsvpHeader">If you're responding for you and a guest (or your family), you'll be able to RSVP for your entire group.</div>
        <div class="inputWrapper" style="min-width:100%">
            <input class="searchBar" type="text" bind:value={name} required/>
            <label>Full Name</label>
        </div>
        <button class="searchButton" type="submit" onclick={handleSubmit}>Find Your Invitation</button>

        {#if error}
            <p class="error">{error}</p>
        {/if}
    {/if}

    {#if currentStage === 2 && userData}
        <h2>RSVP For Your Household</h2>
        {#each $userData.guests as guest}
            <GuestInput bind:guest/>
        {/each}
        <div class="buttonGroup">
            <!--  Todo: Disable the "Next" button unless all guests are non-null     -->
            <button class="nextButton" type="button" onclick={handleBack}>Back</button>
            <button class="nextButton" type="button" onclick={handleNext}>Next</button>
        </div>
    {/if}

    {#if currentStage === 3 && userData}
        <h2>Dinner Selection</h2>
        {#each $userData.guests as guest}
            <DinnerSelect bind:guest/>
        {/each}
        <div class="buttonGroup">
            <button class="nextButton" type="button" onclick={handleBack}>Back</button>
            <button class="nextButton" type="button" onclick={handleNext}>Next</button>
        </div>
    {/if}

    {#if currentStage === 4 && userData}
        <h2>Any dietary restrictions?</h2>
        <textarea class="noteField" bind:value={$userData.notes}></textarea>
        <button class="searchButton" type="submit" onclick={handleSave}>Continue</button>
    {/if}

    {#if currentStage === 5}
        <h2>Your RSVP Has Been Saved!</h2>
        <a href="/" class="searchButton">Back To Home</a>
    {/if}
</div>



<style>
    .noteField {
        font-family: "Times New Roman", serif;
        letter-spacing: 1px;
        font-size: clamp(14px, calc(0.875rem + ((1vw - 4px) * 0.3765)), 19px);
        font-weight: 400;
        background-color: rgb(246, 233, 213) !important;
        border-color: rgb(64, 64, 65) !important;
        color: rgb(64, 64, 65);
        padding: 1.25rem .75rem .25rem .625rem;
        resize: none;
        transition-property: border-color;
        margin: 0;
        height: 140px;
        max-width: 440px;
        width: 80vw;
        text-align: left;
        z-index: 2;
    }

    .buttonGroup {
        margin-top: 5rem;
        display: flex;
        justify-content: center;
        gap: 2rem;
    }

    .searchBar {
        appearance: none;
        background-color: rgb(255, 255, 255);
        border-color: black;
        border-radius: 2px;
        border-style: solid;
        border-width: 1px;
        border-image-repeat: stretch;
        box-shadow: none;
        box-sizing: border-box;
        color: rgba(0, 0, 0, 0.4);
        cursor: text;
        display: block;
        font-weight: 400;
        font-size: 1rem;
        height: 3rem;
        letter-spacing: 1px;
        line-height: 1.5;
        padding-block-start: 15px;
        padding-inline-end: 10px;
        padding-inline-start: 10px;
        padding: 15px .625rem 0;
        position: relative;
        text-align: start;
        text-indent: 0px;
        text-rendering: auto;
        text-size-adjust: 100%;
        touch-action: manipulation;
        transition-delay: 0s;
        transition-duration: 0.2s;
        transition-property: border-color;
        transition-timing-function: ease-out;
        width: 100%;
        word-spacing: 0px;
    }

    .searchButton {
        margin: 20px;
        line-height: 2.25rem;
        max-width: 400px;
        padding-left: 1.25rem;
        padding-right: 1.25rem;
        cursor: pointer;
        padding-bottom: 0px;
        padding-top: 0.125rem;
        text-align: center;
        letter-spacing: 2px;
        color: rgb(17, 71, 57);
        font-size: 19px;
        font-family: "Times New Roman",serif;
        background-color: rgb(128, 171, 127);
        text-transform: uppercase;
        display: block !important;
        width: 100% !important;
        border-radius: 2px;
        transition: background-color 0.1s ease-out, color, border-color;
        border-width: 1px;
        border-style: solid;
        border-image: initial;
        border-color: rgba(0, 0, 0, 0.4);
    }

    .nextButton {
        line-height: 2.25rem;
        max-width: 400px;
        padding-left: 1.25rem;
        padding-right: 1.25rem;
        cursor: pointer;
        padding-bottom: 0px;
        padding-top: 0.125rem;
        text-align: center;
        letter-spacing: 2px;
        color: rgb(17, 71, 57);
        font-size: 19px;
        font-family: "Times New Roman",serif;
        background-color: rgb(128, 171, 127);
        text-transform: uppercase;
        width: 50% !important;
        border-radius: 2px;
        transition: background-color 0.1s ease-out, color, border-color;
        border-width: 1px;
        border-style: solid;
        border-image: initial;
        border-color: rgba(0, 0, 0, 0.4);
    }


</style>