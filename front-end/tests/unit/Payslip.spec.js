import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import PayslipPage from '../../src/components/PayslipPage.vue'

describe('PayslipPage.vue', () => {
    it('Employees payslip page title is being rendered correctly ', () => {
        const wrapper = shallowMount(PayslipPage);
        expect(wrapper.find('#titleId').exists()).to.exist;
        expect(wrapper.find('#titleId').text()).equal('Payslip Page');
    });

    it('Employees payslips are being displayed month wise ', () => {
        const wrapper = shallowMount(PayslipPage);
        expect(wrapper.find('#empTableId').exists()).to.exist;
    });

    it('Employees payslips are being displayed month wise ', () => {
        const wrapper = shallowMount(PayslipPage);
        expect(wrapper.find('#teamButtonId').exists()).to.exist;
        expect(wrapper.find('#teamButtonId').text()).equal('Show Team Payslip');
    });
})