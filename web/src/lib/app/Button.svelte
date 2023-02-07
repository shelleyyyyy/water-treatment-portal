

<script>

export let options;
export let publish;

const click = (c) => {
    publish(options.click)
}

</script>



<div class="bg-black">
    <button on:click={click} class="btn w-full">{options.click}</button>
</div>