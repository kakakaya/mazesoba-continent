import type { Meta, StoryObj } from '@storybook/svelte';
import InputBox from './InputBox.svelte';

// Storybook for InputBox.svelte


const meta = {
    component: InputBox,
    title: 'Components/InputBox',
} satisfies Meta<InputBox>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        count: 0,
    },
};
