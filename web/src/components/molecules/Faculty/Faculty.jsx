import { Card, CardContent, Typography } from "@mui/material";

function Faculty({ faculty }) {
  return (
    <Card>
      <CardContent>
        <Typography
          variant="h6"
          sx={{
            WebkitBoxOrient: "vertical",
            WebkitLineClamp: {
              xs: "2",
              md: "1",
            },
            display: "-webkit-box",
            overflow: "hidden",
          }}
        >
          {faculty.name}
        </Typography>
      </CardContent>
    </Card>
  );
}

export default Faculty;
