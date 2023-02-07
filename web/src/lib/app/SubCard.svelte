<script>

    import { pb } from "$lib/app/pocketbase.js"
	import { onMount } from "svelte"

    export let id;

    $: record = "tmp"

    var title = ""

    onMount(async () => {

        var res = await pb.collection('subsciptions').getOne(id)
        console.log(res)
        title = res.title
        record = res.value

        pb.collection('subsciptions').subscribe(id, function (e) {
            record = e.record.value
        });
    })

</script>



<div class="p-5 bg-secondary rounded-lg">
    <div class="font-bold text-center">
        {title}
    </div>
    <div class="text-center p-3">
        {record}
    </div>
</div>