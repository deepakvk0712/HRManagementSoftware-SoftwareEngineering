import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import resignationPopup from '../../src/Popups/resignationPopup.vue'

describe('resignationPopup.vue', () => {
    it('Renders the HR register employee form heading ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#finCardHead').text()).equal('Resignation Form');
    });

    it('Renders the various input fields in resignation page ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#train-rate').exists()).to.exist;
        expect(wrapper.find('#growth-rating').exists()).to.exist;
        expect(wrapper.find('#recommend-other').exists()).to.exist;
        expect(wrapper.find('#product-rating').exists()).to.exist;
        expect(wrapper.find('#stay-rating').exists()).to.exist;
        expect(wrapper.find('#resignFeedback').exists()).to.exist;
        expect(wrapper.find('#changeMind').exists()).to.exist;
    });

    it('Renders the various checkboxes in resignation page ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#agree-checkbox').exists()).to.exist;
        expect(wrapper.find('#document-checkbox').exists()).to.exist;
        expect(wrapper.find('#terms-checkbox').exists()).to.exist;
    });

    it('Renders the full name and current date input ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#full-name').exists()).to.exist;
        expect(wrapper.find('#filterEndDate').exists()).to.exist;
    });

    it('Renders the cancel button and submit button element ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#finCancelButton').exists()).to.exist;
        expect(wrapper.find('#finSubmitButton').exists()).to.exist;
    });

    it('Renders the text on the cancel and submit button correctly ', () => {
        const wrapper = shallowMount(resignationPopup);
        expect(wrapper.find('#finCancelButton').text()).equal('Cancel');
        expect(wrapper.find('#finSubmitButton').text()).equal('Resign');
    });

})