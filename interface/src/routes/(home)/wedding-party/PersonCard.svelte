<script>
    export let photo = "";
    export let name = "";
    export let role = "";
    export let relation = ""
</script>


<div class="personCard">
    <div class="photo-container">
        <img src={photo} alt={name} class="photo"/>
    </div>
    <div class="info">
        <p class="name">{name}</p>
        <p class="details">{role}, {relation}</p>
    </div>
</div>

<style>
    .personCard {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-bottom: 2rem;
    }

    .photo-container {
        width: 20vw;
        height: 20vw;
        border-radius: 50%;
        overflow: hidden;
        margin-bottom: 1rem;
    }

    .photo {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .info {
        text-align: center;
        color: rgba(0, 0, 0, 0.4);
        font-size: clamp(10px, calc(0.625rem + ((1vw - 4px) * 0.6024)), 18px);;
        font-weight: 400;
        text-transform: uppercase;
        letter-spacing: .42vw;
        max-width: 80%;
    }
</style>

