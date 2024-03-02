import { Card, Typography } from "@mui/material";

function Faculty({ faculty }) {
  return (
    <Card>
      <Typography>{faculty.name}</Typography>
    </Card>
  );
}

export default Faculty;
