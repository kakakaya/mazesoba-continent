import type { Meta, StoryObj } from '@storybook/svelte';
import StatusBar from './StatusBar.svelte';

const meta = {
    component: StatusBar,
    title: 'Components/StatusBar',
} satisfies Meta<StatusBar>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        helpMessage: 'This is a help message',
        postFooter: '#test #storybook',
        charCount: 0,
        maxCount: 300,
    },
};


export const OverLimit: Story = {
    args: {
        helpMessage: 'This is a help message',
        postFooter: '#test #storybook',
        charCount: 400,
        maxCount: 300,
    },
};

export const LongFooter: Story = {
    args: {
        helpMessage: 'Super Long Message is Here',
        postFooter: '#Lorem #ipsum #dolor #sit #amet, #consectetur #adipiscing #elit, #sed #do #eiusmod #tempor #incididunt #ut #labore #et #dolore #magna #aliqua. #Ut #enim #ad #minim #veniam, #quis #nostrud #exercitation #ullamco #laboris #nisi #ut #aliquip #ex #ea #commodo #consequat. #Duis #aute #irure #dolor #in #reprehenderit #in #voluptate #velit #esse #cillum #dolore #eu #fugiat #nulla #pariatur. #Excepteur #sint #occaecat #cupidatat #non #proident, #sunt #in #culpa #qui #officia #deserunt #mollit #anim #id #est #laborum.',
        charCount: 0,
        maxCount: 300,
    },
};
