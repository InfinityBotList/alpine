<script>
    import '@svelteness/kit-docs/client/polyfills/index.js';
    import '@svelteness/kit-docs/client/styles/normalize.css';
    import '@svelteness/kit-docs/client/styles/theme.css';
    import '@svelteness/kit-docs/client/styles/vars.css';
    
    import { page } from '$app/stores';
  
    import {
      createKitDocsLoader,
      KitDocs,
      KitDocsLayout,
      createSidebarContext,
    } from '@svelteness/kit-docs';
  
    export let data;
  
    let { meta, sidebar } = data;
    $: ({ meta, sidebar } = data);
  
    const { activeCategory } = createSidebarContext(sidebar);
  
    $: category = $activeCategory ? `${$activeCategory}: ` : '';
    $: title = meta ? `${category}${meta.title} | Svelte` : null;
    $: description = meta?.description;
  </script>
    
  <KitDocs {meta}>
    <KitDocsLayout navbar={false} {sidebar} search>  
      <input type="text" placeholder="Search documentation" slot="search" />  
      <slot />
    </KitDocsLayout>
  </KitDocs>