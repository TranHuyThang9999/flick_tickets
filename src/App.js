import React from "react";

import HomeaAdmin from "./dashboard/HomeaAdmin";
import CreateAccountStaff from "./dashboard/CreateAccountStaff";
import GetAllStaff from "./dashboard/GetAllStaff";
import SelectedSeat from './common/cinemas/SelectedSeat';
import AppSquare from "./common/test/test";
import DisplaySelectedSeat from "./dashboard/DisplaySelectedSeat";
import GetAllTickets from "./dashboard/GetAllTickets";
import GetTicketById from "./dashboard/GetTicketById";
import DetailedShowSchedule from "./common/Showtimes/DetailedShowSchedule";
import GetListFileByTicketId from "./dashboard/GetListFileByTicketId";
import PageForUser from "./Home/Page/PageForUser";
import GetListTicket from "./Home/GetListTickets/GetListTicket";

export default function App() {
  return (
    <div>
      {/* <HomeaAdmin /> */}
      {/* <CreateAccountStaff/> */}
      {/* <GetAllStaff/> */}
      {/* <AppTest/> */}
      {/* <SelectedSeat SelectedSeatGetFormApi={[1,45,23,90]} numSquares={90} heightContainerUseSaveData={500} widthContainerUseSavedate={900}/> */}
      {/* <AppSquare/> */}
      {/* <DisplaySelectedSeat/> */}
      {/* <GetAllTickets/> */}
      {/* <GetTicketById id={7133536}/> */}
      {/* <DetailedShowSchedule id={7133536}/> */}
      {/* <GetListFileByTicketId id={7133536}/> */}
      {/* <PageForUser/> */}
      <GetListTicket/>
    </div>
  );
}
