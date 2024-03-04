import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';

const PieChart = ({ data, labels, colors }) => {
  const chartRef = useRef(null);
  let myChart = null;

  useEffect(() => {
    if (chartRef.current && myChart) {
      myChart.destroy(); // Si el gráfico ya existe, destrúyelo
    }

    const ctx = chartRef.current.getContext('2d');
    myChart = new Chart(ctx, {
      type: 'pie',
      data: {
        labels: labels,
        datasets: [{
          data: data,
          backgroundColor: colors,
          borderColor: colors,
          hoverBackgroundColor: colors, //['#269fbd', '#87f4e8'],
            borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false
      }
    });
    return () => {
      if (myChart) {
        myChart.destroy(); // Asegúrate de destruir el gráfico al desmontar el componente
      }
    };
  }, [data, labels, colors]);

  return (
    <div>
      <canvas ref={chartRef}></canvas>
    </div>
  );
};

export default PieChart;
