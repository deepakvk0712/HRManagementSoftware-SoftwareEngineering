import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import AboutView from '../../src/views/AboutView.vue'

describe('AboutView.vue', () => {
    it('Renders the tabs in the about view page ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#tabsId').exists()).to.exist;
    });

    it('Printing the names of the tabs correctly ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#tabsId').exists()).to.exist;
    });

    it('Renders the about us heading correctly ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#AboutUsHeading').text()).equal('Meet Axle United!');
        expect(wrapper.find('#textAboutUs').exists()).to.exist;
    });

    it('Renders the team information correctly ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#teamAboutUs').text()).equal('Meet The Team');
        expect(wrapper.find('#accomplishmentsTeams').exists()).to.exist;
    });

    it('Renders the company policy correctly ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#companyPolicy').text()).equal('Company Policy');
        expect(wrapper.find('#companyPolicyList').exists()).to.exist;
    });

    it('Renders the employee policy correctly ', () => {
        const wrapper = shallowMount(AboutView);
        expect(wrapper.find('#employeePolicy').text()).equal('Employee Policy');
        expect(wrapper.find('#employeePolicyList').exists()).to.exist;
    });

    // it('Renders the core values correctly ', () => {
    //     const wrapper = shallowMount(AboutView);
    //     expect(wrapper.find('#employeePolicy').text()).equal('Employee Policy');
    //     expect(wrapper.find('#employeePolicyList').exists()).to.exist;
    // });

    // it('Renders the Mission Statement correctly ', () => {
    //     const wrapper = shallowMount(AboutView);
    //     expect(wrapper.find('#employeePolicy').text()).equal('Employee Policy');
    //     expect(wrapper.find('#employeePolicyList').exists()).to.exist;
    // });
})