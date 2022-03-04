// import { expect } from 'chai'
// import { shallowMount } from '@vue/test-utils'
// // import { mount } from "@vue/test-utils";
// import LandingPage from '../../src/components/LandingPage.vue'
// // import FinancialFormPopup from '../../src/Popups/FinancialFormPopup.vue'

// describe('LandingPage.vue', () => {
//   it('Renders required forms and quick links headings ', () => {
//     // const msg = 'new message'
//     // const wrapper = shallowMount(LandingPage, {
//     //   propsData: { msg }
//     // })
//     // expect(wrapper.text()).to.include(msg)

//     const wrapper = shallowMount(LandingPage);
//     expect(wrapper.find('#quicklink1').text()).equal('Required Forms');
//     expect(wrapper.find('#quicklink2').text()).equal('Quick Links');
//   });

//   it('Renders the heading of each required form ', () => {
//     const wrapper = shallowMount(LandingPage);
//     expect(wrapper.find('#regEmpHead').text()).equal('Register Employee');
//     expect(wrapper.find('#onboardForm').text()).equal('Onboarding Form');
//     expect(wrapper.find('#finForm').text()).equal('Financial Form');
//   });

//   it('Renders the sub-heading of each required form ', () => {
//     const wrapper = shallowMount(LandingPage);
//     expect(wrapper.find('#subRegEmpHead').text()).equal('Register new employee');
//     expect(wrapper.find('#subOnboardForm').text()).equal('Employee Input Needed');
//     expect(wrapper.find('#subFinForm').text()).equal('Banking Information');
//   });

//   // it('Renders the text of the financial form popup button ', () => {
//   //   const wrapper = mount(LandingPage);
//   //   expect(wrapper.find('#finButtonText').text()).equal('Fill Form');
//   //   // expect(wrapper.find('#subOnboardForm').text()).equal('Employee Input Needed');
//   //   // expect(wrapper.find('#subFinForm').text()).equal('Banking Information');
//   // });
// })
