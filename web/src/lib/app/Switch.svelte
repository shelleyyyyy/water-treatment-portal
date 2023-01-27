<script>

export let options;
export let publish;

let checked = true;

const flip = (c) => {
    if(c){
        publish(options.on)
    } else{
        publish(options.off)
    }
}


</script>

<div class="grid grid-cols-1 ">
    <div class="flex justify-around ">
        <div class="">
            {options.on}
        </div>
        <div class="">
            {options.off}
        </div>
    </div>
    <div class="p-5">
        <input type="checkbox" class="toggle" on:change={flip(checked)} bind:checked />
    </div>
</div>