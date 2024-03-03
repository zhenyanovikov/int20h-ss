import { Grid } from "@mui/material";
import { PieChart } from "@mui/x-charts/PieChart";
import GroupTemplate from "../GroupTemplate/GroupTemplate";

function GroupAnalyticsTemplate() {
  return (
    <GroupTemplate>
      <Grid container spacing={2}>
        <Grid item xs={12} md={6}>
          <PieChart
            series={[
              {
                data: [
                  { id: 0, value: 3, label: "A" },
                  { id: 1, value: 5, label: "B" },
                  { id: 2, value: 7, label: "C" },
                  { id: 3, value: 11, label: "D" },
                  { id: 4, value: 6, label: "E" },
                  { id: 5, value: 1, label: "F" },
                ],
              },
            ]}
            width={400}
            height={200}
          />
        </Grid>
      </Grid>
    </GroupTemplate>
  );
}

export default GroupAnalyticsTemplate;
