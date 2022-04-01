// import { expect } from 'chai'
// import { shallowMount } from '@vue/test-utils'
// import Attendance from '../../src/components/AttendancePage.vue'

// describe('Attendance.vue', () => {
//     it('Renders the Attendance page heading ', () => {
//         const wrapper = shallowMount(Attendance);
//         expect(wrapper.find('#headingSubmit').text()).equal('Submit Attendance');
//     });

//     it('Renders the calendar to take the date ', () => {
//         const wrapper = shallowMount(AttendencePage);
//         expect(wrapper.find('#datePickingSubmit').exists()).to.exist;
//     });

//     it('Renders the time components to take the start and end time ', () => {
//         const wrapper = shallowMount(AttendencePage);
//         expect(wrapper.find('#startTimeSubmit').exists()).to.exist;
//         expect(wrapper.find('#endTimeSubmit').exists()).to.exist;
//     });

//     it('Renders the Reset and Submit buttons and the text on each button ', () => {
//         const wrapper = shallowMount(AttendencePage);
//         expect(wrapper.find('#resetButton').exists()).to.exist;
//         expect(wrapper.find('#submitButton').exists()).to.exist;
        
//         expect(wrapper.find('#resetButton').text()).equal('Reset');
//         expect(wrapper.find('#submitButton').text()).equal('Submit');
//     });

//     it('Renders the Start/End date picker for filtering and binds the value chosen ', () => {
//         const wrapper = shallowMount(AttendencePage);
//         expect(wrapper.find('#filterStartDate').exists()).to.exist;
//         expect(wrapper.find('#filterEndDate').exists()).to.exist;
//     });

//     it('Renders the Reset and Submit buttons and the text on each button ', () => {
//         const wrapper = shallowMount(AttendencePage);
//         expect(wrapper.find('#filterButton').exists()).to.exist;

//         expect(wrapper.find('#filterButton').text()).equal('Filter');
//     });
// })