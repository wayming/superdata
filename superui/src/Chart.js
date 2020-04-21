import React from 'react';
import './App.css';
import '../node_modules/nvd3/build/nv.d3.css';
import NVD3Chart from "react-nvd3";
import d3 from "d3";
import { useState } from 'react';

function Chart() {
  const [units, setUnits] = useState(false);
  fetch('http://localhost:8001')
    .then(response => {
      return response.json();
    })
    .then(data => {
      // console.log(data)
      setUnits(data);
    });
  return (
    <div className="Chart">
        {
          React.createElement(NVD3Chart, {
            xAxis: {
              // tickFormat: function(d){ return Date(d); },
              tickFormat: function(d){ return d3.time.format('%x')(new Date(d)); },
              axisLabel: 'Period'
            },
            yAxis: {
              // tickFormat: function(d) {return parseFloat(d).toFixed(6); }
              tickFormat: d3.format(',.1%')
            },
            type:'cumulativeLineChart',
            datum: units  ,
            x: function(d) {
              return new Date(d['vdate']);
            },
            y: function(d) {
              return d['value'];
            },
            height: 600
          })
        }
      <h1> { JSON.stringify(units) } } </h1>

    </div>
  );
}

export default Chart;
