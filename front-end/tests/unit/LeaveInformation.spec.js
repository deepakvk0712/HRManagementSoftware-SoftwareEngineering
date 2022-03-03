import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import ApplyLeavePopup from '../../src/Popups/ApplyLeavePopup.vue'

describe('HRRegisterEmployee.vue', () => {
    it('Renders the leave application popup form heading ', () => {
        const wrapper = shallowMount(ApplyLeavePopup);
        expect(wrapper.find('#leaveApplyHead').text()).equal('Leave Application');
    });

    it('Renders the various input fields in leave application popup ', () => {
        const wrapper = shallowMount(ApplyLeavePopup);
        expect(wrapper.find('#leaveApplySubjectInput').exists()).to.exist;
        expect(wrapper.find('#leaveApplyReasonInput').exists()).to.exist;
        expect(wrapper.find('#leaveApplyLeaveTypeInput').exists()).to.exist;
        expect(wrapper.find('#leaveApplyEndDateInput').exists()).to.exist;
        // expect(wrapper.find('#regEmpTypeInput').exists()).to.exist;
    });

    it('Renders the Cancel and Submit buttons and button text in leave application popup ', () => {
        const wrapper = shallowMount(ApplyLeavePopup);
        expect(wrapper.find('#leaveApplyCancelInput').exists()).to.exist;
        expect(wrapper.find('#leaveApplySubmitInput').exists()).to.exist;
        
        expect(wrapper.find('#leaveApplyCancelInput').text()).equal('Cancel');
        expect(wrapper.find('#leaveApplySubmitInput').text()).equal('Apply');
    });
})