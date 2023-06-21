<script lang="ts">
    import '@svelteness/kit-docs/client/polyfills/index.js';
    import '@svelteness/kit-docs/client/styles/vars.css';
    import { onMount } from 'svelte';

    import {
      createKitDocsLoader,
      KitDocs,
      KitDocsLayout,
      createSidebarContext,
	  Button,
    } from '@svelteness/kit-docs';
	import { getLinks } from '$lib/corelib/links.js';
  
    export let data;
  
    let { meta, sidebar } = data;
    $: ({ meta, sidebar } = data);
  
    const { activeCategory } = createSidebarContext(sidebar);
  
    $: category = $activeCategory ? `${$activeCategory}: ` : '';
    $: title = meta ? `${category}${meta.title} | Svelte` : null;
    $: description = meta?.description;

    let links: {
      title: string;
      slug: string;
    }[] = []
    onMount(async () => {
      let linkFound = await getLinks()

      if (linkFound) {
        links = linkFound.map(link => {
          return {
            title: link?.Title,
            slug: link?.Href
          }
        })
      }
    })
</script>
    
  <KitDocs {meta}>
    <KitDocsLayout navbar={{
      links: [
        { slug: "/", title: "Back Home" },
        ...links
      ]
    }} {sidebar}>  
      <div slot="navbar-left">
        <div class="logo">
          <Button href="/">
            <img src="https://cdn.infinitybots.gg/images/png/Infinity.png" class="mr-3 h-6 sm:h-9" alt="IBL Logo" />
            <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Infinity Docs</span>
          </Button>
        </div>
      </div>
      <slot />
    </KitDocsLayout>
  </KitDocs>