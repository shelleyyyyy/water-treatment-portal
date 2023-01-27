import { error } from '@sveltejs/kit';
 
/** @type {import('./$types').PageLoad} */
export function load({ params }) {

    return {
        title: params.slug,
        content: 'Welcome to our blog. Lorem ipsum dolor sit amet...'
    };
 
    throw error(404, 'Not found');
}