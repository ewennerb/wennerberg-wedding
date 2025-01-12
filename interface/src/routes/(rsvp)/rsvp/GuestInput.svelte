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
    <div class="radioContainer">
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

    @media (max-width: 400px) {
        label.selected {
            padding: 0.5rem;
        }
        label.deselected {
            padding: 0.5rem;
        }
    }

    input[type="radio"] {
        display: none;
    }

    .radioContainer {
        display: flex;
        width: 50%;
        margin-bottom: 0px;
        justify-content: center;
    }

    span {
        font-size: clamp(14px, calc(0.875rem + ((1vw - 4px) * 0.4518)), 20px);
    }

</style>