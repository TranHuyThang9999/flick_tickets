import React, { useState, useEffect } from "react";
import { DatePicker, Button } from "antd";
import moment from "moment";

const App = () => {
  const [timestampsList, setTimestampsList] = useState([]);
  const [, setSavedTimestampsList] = useState([]);

  const handleDateChange = (date, dateString) => {
    if (date && moment(date).isValid()) {
      setTimestampsList((prevTimestampsList) => [
        ...prevTimestampsList,
        moment(dateString).unix(),
      ]);
    }
  };

  useEffect(() => {
    const updatedTimestampsList = timestampsList.map((timestamp) =>
      moment.unix(timestamp)
    );
    setTimestampsList(updatedTimestampsList);
  }, []);

  const handleSaveDates = () => {
    setSavedTimestampsList(timestampsList);
    console.log(timestampsList);
  };

  return (
    <div className="App">
      <DatePicker
        showTime
        onChange={handleDateChange}
        picker="datetime"
        size="small"
      />
      <Button onClick={handleSaveDates}>Save Dates</Button>

      <div>
        <h3>List of Selected Dates:</h3>
        {timestampsList.map((timestamp, index) => {
          const date = moment.unix(timestamp);
          return (
            <div key={index}>
              <p>Timestamp: {timestamp}</p>
              <p>
                Hour: {date.format("HH")}, Day: {date.format("DD")}, Month:{" "}
                {date.format("MM")}, Year: {date.format("YYYY")}
              </p>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default App;
