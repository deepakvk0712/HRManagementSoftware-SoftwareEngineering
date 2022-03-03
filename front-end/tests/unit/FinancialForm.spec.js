import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import FinancialFormPopup from '../../src/Popups/FinancialFormPopup.vue'

describe('FinancialFormPopup.vue', () => {
    it('Renders the financial form heading ', () => {
        const wrapper = shallowMount(FinancialFormPopup);
        expect(wrapper.find('#finCardHead').text()).equal('Banking Information');
    });

    it('Renders the 3 input fields - bank name, routing number and account number ', () => {
        const wrapper = shallowMount(FinancialFormPopup);
        expect(wrapper.find('#bankNameInput').exists()).to.exist;
        expect(wrapper.find('#routingNumberInput').exists()).to.exist;
        expect(wrapper.find('#accountNumberInput').exists()).to.exist;
    });

    it('Renders the Cancel and Submit buttons and button text ', () => {
        const wrapper = shallowMount(FinancialFormPopup);
        expect(wrapper.find('#finCancelButton').exists()).to.exist;
        expect(wrapper.find('#finSubmitButton').exists()).to.exist;
        
        expect(wrapper.find('#finCancelButton').text()).equal('Cancel');
        expect(wrapper.find('#finSubmitButton').text()).equal('Submit');
    });
})