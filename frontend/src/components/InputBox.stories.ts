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
        value: 'Input Box',
        placeholder: 'Enter text here',
    },
};

export const Init: Story = {
    args: {
        value: '',
        placeholder: 'どう最近？',
    },
};


export const WithEmoji: Story = {
    args: {
        value: '🙂‍↔️は１文字で表示',
        placeholder: '',
    },
};