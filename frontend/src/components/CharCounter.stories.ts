import type { Meta, StoryObj } from '@storybook/svelte';
import CharCounter from './CharCounter.svelte';

const meta = {
    component: CharCounter,
    title: 'Components/CharCounter',
} satisfies Meta<CharCounter>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        count: 0,
    },
};


export const Maximum: Story = {
    args: {
        count: 300,
    },
};

export const OverLimit: Story = {
    args: {
        count: 301,
    },
};

export const Hidden: Story = {
    args: {
        hidden: true,
        count: 0,
    },
};