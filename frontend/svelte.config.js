import sveltePreprocess from 'svelte-preprocess'
import {mdsvex} from 'mdsvex'

export default {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: [
    sveltePreprocess(),
    mdsvex({
      extensions: ['.svx','.md'],
    })
  ],

  extensions: ['.svelte', '.svx', '.md'],
  
}
