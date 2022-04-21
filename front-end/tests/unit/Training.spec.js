import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import Training from '../../src/components/Training.vue'

describe('Training.vue', () => {
    it('The training is rendered properly ', () => {
        const wrapper = shallowMount(Training);
        expect(wrapper.find('#trainingID').exists()).to.exist;
        expect(wrapper.find('#trainingID').text()).equal('Welcome to the Training Session');
    });

    it('Training page wordings are proper ', () => {
        const wrapper = shallowMount(Training);
        expect(wrapper.find('#tID').exists()).to.exist;
        expect(wrapper.find('#tID').text()).equal('Please watch this video and answer the below questions');
    });

    it('If the training instructions are rendered ', () => {
        const wrapper = shallowMount(Training);
        expect(wrapper.find('#ansID').exists()).to.exist;
        expect(wrapper.find('#ansID').text()).equal('Get 7/10 correct answers to complete training');
    });

    it('The training page correctly displays the expected sentences ', () => {
        const wrapper = shallowMount(Training);
        expect(wrapper.find('#aID').exists()).to.exist;
        expect(wrapper.find('#aID').text()).equal('All the Best!');
    });

    // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });

        // it('If the training score is correctly displayed ', () => {
    //     const wrapper = shallowMount(Training);
    //     expect(wrapper.find('#cID').exists()).to.exist;
    //     expect(wrapper.find('#cID').text()).equal('Thanks for completing the test! Your score is');
    // });


})