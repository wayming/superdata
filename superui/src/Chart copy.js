import React from 'react';
import logo from './logo.svg';
import './App.css';
import '../node_modules/nvd3/build/nv.d3.css';
import NVD3Chart from "react-nvd3";
import d3 from "d3";
import { useState } from 'react';

function Chart() {
  const [units, setUnits] = useState(false);
  console.log("Chart()")
  fetch('http://api:3001')
    .then(response => {
      console.log(response)
      return response.text();
    })
    .then(data => {
      console.log(data)
      setUnits(data);
    });
  return (
    <div className="Chart">
      <p>{units}</p>
        {
          React.createElement(NVD3Chart, {
            // xAxis: {
            //   tickFormat: function(d){ return d; },
            //   axisLabel: 'Period'
            // },
            // yAxis: {
            //   tickFormat: function(d) {return parseFloat(d).toFixed(2); }
            // },
            type:'cumulativeLineChart',
            datum: {units},
            x: 'vdate',
            y: 'value',
            height: 600
          })
        }
    </div>
  );
}

export default Chart;
