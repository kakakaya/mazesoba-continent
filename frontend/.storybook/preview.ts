import type { Preview } from "@storybook/svelte";
import "../src/style.css";
import "../src/storybook.style.css";

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;
