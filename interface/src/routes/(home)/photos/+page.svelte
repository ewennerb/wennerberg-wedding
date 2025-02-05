<script>
    import { onMount } from "svelte";
    import PhotoSwipeLightbox from 'photoswipe/lightbox';
    import 'photoswipe/style.css';
    import Footer from "../Footer.svelte";

    onMount(() => {
        let lightbox = new PhotoSwipeLightbox({
            gallery: '#engagement-gallery',
            children: 'a',
            pswpModule: () => import('photoswipe')
        });
        lightbox.init();})
</script>


<div class="contentContainer">
    <div class="pswp-gallery" id="engagement-gallery">
        <a
            href="/img/gallery/bridge-1.jpg"
            data-pswp-width="1024"
            data-pswp-height="683"
            target="_blank">
            <img src="/img/gallery/bridge-1.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/candy-store-1.jpg"
            data-pswp-width="1024"
            data-pswp-height="683"
            target="_blank">
            <img src="/img/gallery/candy-store-1.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/dip.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/dip.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/gazebo.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/gazebo.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/ivy-wall-2.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/ivy-wall-2.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/kiss-1.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/kiss-1.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/walk-towards.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/walk-towards.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/wall-tree.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/wall-tree.jpg" alt="agga"/>
        </a>
        <a
            href="/img/gallery/waterfall-1.jpg"
            data-pswp-width="1024"
            data-pswp-height="1536"
            target="_blank">
            <img src="/img/gallery/waterfall-1.jpg" alt="agga"/>
        </a>
    </div>
    <Footer/>
</div>


<style>

    img {
        position: relative;
        inset: 0;
        width: 25vw;
        height: 25vw;
        object-fit: cover;
    }

</style>