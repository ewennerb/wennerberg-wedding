<script>
    export let guest;

    const updateGuest = (field, value) => {
        guest[field] = value;
    }

    const updateAttendance = (val) => {
        guest.attending = val;
        if (val === "false"){
            guest.dinner_choice = "";
        }
    }

</script>


<div class="rsvpRow">
    <div class="rsvpRow nameLabel">{guest.name}</div>
    <div class="rsvpRow buttonContainer">
        <label class:selected={guest.attending === "true"} class:deselected={guest.attending !== "true"}>
            <input type="radio" bind:group={guest.attending} value="true" on:change={() => updateAttendance("true")}/>
            <span>{guest.attending === "true" ? "Accepted" : "Accept" }</span>
        </label>

        <label class:selected={guest.attending === "false"} class:deselected={guest.attending !== "false"}>
            <input type="radio" bind:group={guest.attending} value="false" on:change={() => updateAttendance("false")}/>
            <span>{guest.attending === "false" ? "Declined" : "Decline" }</span>
        </label>
    </div>
</div>


<style>


    label.deselected {
        margin-left: 0.5rem;
        margin-right: 0.5rem;
        min-width: 70px;
        display: inline-block;
        background: #f6e9d5;
        border: 1px solid rgba(0, 0, 0, 0.4);
        padding: 0.5rem 1.25rem;
        border-radius: 5px;
        cursor: pointer;
        color: rgba(0, 0, 0, 0.4);
        transition: box-shadow 400ms ease, background 400ms ease, color 400ms ease;
        text-align: center;
    }

    label.selected {
        margin-left: 0.5rem;
        margin-right: 0.5rem;
        min-width: 70px;
        display: inline-block;
        background: #80ab7f;
        border: 1px solid rgb(128, 171, 127);
        padding: 0.5rem 1.25rem;
        border-radius: 5px;
        cursor: pointer;
        color: rgb(246, 233, 213);
        transition: box-shadow 400ms ease, background 400ms ease, color 400ms ease;
        text-align: center;
    }

    input[type="radio"] {
        display: none;
    }


    .rsvpRow {
        height: 80%;
        width: 480px;
        padding: 5px 20px 0px;
        text-align: left;
        z-index: 2;
        display: flex;
        -webkit-box-pack: justify;
        justify-content: space-between;
        -webkit-box-align: center;
        align-items: center;
        margin-bottom: 24px;
    }

    .rsvpRow .nameLabel {
        color: rgba(0, 0, 0, 0.4);
        font-weight: 400;
        display: flex;
        font-size: 20px;
        margin-bottom: 0px;
    }

    .rsvpRow .buttonContainer {
        display: flex;
        width: 50%;
        margin-bottom: 0px
    }

    .knotRow .inputStuff .buttonWrapper {
        height: 40px;
        font-size: 16px;
        border-radius: 2px;
        align-items: center;
        justify-content: center;
        display: flex;
        opacity: 1;
    }

    /*.acceptButton {*/
    /*    height: 40px;*/
    /*    width: 124px;*/
    /*    font-size: 16px;*/
    /*    border-radius: 2px;*/
    /*    background-color: rgb(246, 233, 213);*/
    /*    border: 2px solid rgba(0, 0, 0, 0.4);*/
    /*    color: rgba(0, 0, 0, 0.4);*/
    /*    opacity: 0.5;*/
    /*    margin-left: 8px;*/
    /*    margin-right: 12px;*/
    /*    touch-action: manipulation;*/
    /*}*/

    /*.rsvpField {*/
    /*    margin-right: 1rem;*/
    /*    background-color: rgb(255, 255, 255);*/
    /*    border-color: rgba(0, 0, 0, 0.4);*/
    /*    border-radius: 3px;*/
    /*    border-style: solid;*/
    /*    border-width: 1px;*/
    /*    color: rgba(0, 0, 0, 0.4);*/
    /*    cursor: text;*/
    /*    display: block;*/
    /*    font-weight: 400;*/
    /*    height: 90%;*/
    /*    letter-spacing: 1px;*/
    /*    line-height: 24px;*/
    /*    transition-delay: 0s;*/
    /*    transition-duration: 0.2s;*/
    /*    transition-property: border-color;*/
    /*    transition-timing-function: ease-out;*/
    /*    width: 5rem;*/
    /*    word-spacing: 0px;*/
    /*    padding-block-start: 15px;*/
    /*    padding-inline-end: 10px;*/
    /*    padding-inline-start: 10px;*/
    /*    padding-left: 10px;*/
    /*    padding-right: 10px;*/
    /*    padding-top: 0px;*/
    /*    position: relative;*/
    /*    text-align: start;*/
    /*}*/

    /*.guestDropdown {*/
    /*    appearance: none;*/
    /*    background-color: rgb(255, 255, 255);*/
    /*    border-color: black;*/
    /*    border-radius: 2px;*/
    /*    border-style: solid;*/
    /*    border-width: 1px;*/
    /*    border-image-repeat: stretch;*/
    /*    box-shadow: none;*/
    /*    box-sizing: border-box;*/
    /*    color: rgba(0, 0, 0, 0.4);*/
    /*    cursor: text;*/
    /*    display: block;*/
    /*    font-weight: 400;*/
    /*    height: 90%;*/
    /*    letter-spacing: 1px;*/
    /*    line-height: 24px;*/
    /*    padding-block-start: 15px;*/
    /*    padding-inline-end: 10px;*/
    /*    padding-inline-start: 10px;*/
    /*    padding-left: 10px;*/
    /*    padding-right: 10px;*/
    /*    padding-top: 0px;*/
    /*    position: relative;*/
    /*    text-align: start;*/
    /*    text-indent: 0px;*/
    /*    text-rendering: auto;*/
    /*    text-size-adjust: 100%;*/
    /*    touch-action: manipulation;*/
    /*    transition-delay: 0s;*/
    /*    transition-duration: 0.2s;*/
    /*    transition-property: border-color;*/
    /*    transition-timing-function: ease-out;*/
    /*    width: 15rem;*/
    /*    word-spacing: 0px;*/
    /*}*/

</style>