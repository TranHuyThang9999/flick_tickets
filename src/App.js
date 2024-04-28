import React, { useState } from "react";

import HomeaAdmin from "./dashboard/HomeaAdmin";
import CreateAccountStaff from "./dashboard/CreateAccountStaff";
import GetAllStaff from "./dashboard/GetAllStaff";
import SelectedSeat from "./common/cinemas/SelectedSeat";
import AppSquare from "./common/test/test";
import DisplaySelectedSeat from "./dashboard/DisplaySelectedSeat";
import GetListFileByTicketId from "./dashboard/GetListFileByTicketId";
import PageForUser from "./Home/Page/PageForUser";
import GetTicketById from "./Home/Tickets/DetailTicketById";
import ListFileByTicketId from "./Home/Tickets/ListFileByTicketId";
import GetListTicket from "./Home/Tickets/GetListTicket";
import AdminUploadTickets from "./dashboard/AdminUploadTickets";
import TestDetailedShowSchedule from "./common/test/tesst1";
import SendEmail from "./common/ProcessEmail/SendEmail";
import VerifyEmail from "./common/ProcessEmail/VerifyEmail";
import FormCheckOtpEmail from "./common/ProcessEmail/FormCheckOtpEmail";
import DemoSend from "./common/test/test";
import Cookies from "js-cookie"; // Import thư viện js-cookie
import CinemasAdd from "./common/cinemas/CinemasAdd";
import MovieAdd from "./common/MovieTypes/MovieUpload";
import GetAllTicketsForAdmin from "./dashboard/GetAllTickets";
import GetTicketByIdOnForm from './dashboard/GetTicketByIdOnForm';

export default function App() {
  return (
    <div>
      {/* <HomeaAdmin /> */}
      {/* <CreateAccountStaff/> */}
      {/* <GetAllStaff/> */}
      {/* <AppTest/> */}
      {/* <SelectedSeat SelectedSeatGetFormApi={[1,12,13,14,15]} numSquares={100} heightContainerUseSaveData={500} widthContainerUseSavedate={900} onCreate={setSelectPopChid}/> */}
      {/* <AppSquare/> */}
      {/* <DisplaySelectedSeat/> */}
      {/* <GetTicketById id={7133536}/> */}
      {/* <GetListFileByTicketId id={7133536}/> */}
      {/* <PageForUser/> */}
      {/* <GetListTicket/> */}{/**/}
      {/* <ListFileByTicketId id={1485296}/> */}
      {/* <AdminUploadTickets/> */}
      {/* <TestDetailedShowSchedule/> */}
      {/* <SendEmail/>
      <VerifyEmail/> */}
      {/* <FormCheckOtpEmail/> */}
      {/* <DemoSend/> */}
      {/* <GetTicketById id={1485296} /> */}
      {/* <CinemasAdd/> */}
      {/* <MovieAdd/> */}
      {/* <GetAllTicketsForAdmin/> */}
      <GetTicketByIdOnForm id={3244826}/>
    </div>
  );
}
