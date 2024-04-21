import App from "../App.svelte";
import type { Meta, StoryObj } from '@storybook/svelte';

const meta = {
    component: App,
    title: 'App',
    parameters: {
        layout: 'fullscreen',
    },
} satisfies Meta<App>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};