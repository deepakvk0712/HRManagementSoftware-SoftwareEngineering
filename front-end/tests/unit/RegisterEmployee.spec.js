import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import HRRegisterEmployee from '../../src/Popups/HRRegisterEmployee.vue'

describe('HRRegisterEmployee.vue', () => {
    it('Renders the HR register employee form heading ', () => {
        const wrapper = shallowMount(HRRegisterEmployee);
        expect(wrapper.find('#regEmployeeHead').text()).equal('HR Register Employee');
    });

    it('Renders the various input fields ', () => {
        const wrapper = shallowMount(HRRegisterEmployee);
        expect(wrapper.find('#regEmpNameInput').exists()).to.exist;
        expect(wrapper.find('#regEmpLastNameInput').exists()).to.exist;
        expect(wrapper.find('#regEmpBusUnitInput').exists()).to.exist;
        expect(wrapper.find('#regEmpTitleInput').exists()).to.exist;
        expect(wrapper.find('#regEmpTypeInput').exists()).to.exist;
        expect(wrapper.find('#regEmpManagerIdInput').exists()).to.exist;
        expect(wrapper.find('#regEmpGradeInput').exists()).to.exist;
        expect(wrapper.find('#regEmpLocationInput').exists()).to.exist;
        expect(wrapper.find('#regEmpPersEmailInput').exists()).to.exist;
        expect(wrapper.find('#regEmpCountryInput').exists()).to.exist;
    });

    it('Renders the Cancel and Submit buttons and button text ', () => {
        const wrapper = shallowMount(HRRegisterEmployee);
        expect(wrapper.find('#regEmpCancelButton').exists()).to.exist;
        expect(wrapper.find('#regEmpSubmitButton').exists()).to.exist;
        
        expect(wrapper.find('#regEmpCancelButton').text()).equal('Cancel');
        expect(wrapper.find('#regEmpSubmitButton').text()).equal('Submit');
    });
})