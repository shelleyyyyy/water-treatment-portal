<script>

    export let title;
    export let description;
    export let link;
    export let img;

    import Icon from "$lib/img/grit.png"

</script>


<div class="card lg:card-side bg-base-100 shadow-xl">
    <figure><img src={Icon} alt={img}/></figure>
    <div class="card-body">
        <h2 class="card-title">{title}!</h2>
        <p>{description}</p>
        <div class="card-actions justify-end">
            <a href={"/app/" + link} class="btn btn-primary">View</a>
        </div>
    </div>
</div>
