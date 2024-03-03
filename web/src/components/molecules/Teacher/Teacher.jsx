import { Card, CardContent } from "@mui/material";
import UserPreview from "../UserPreview/UserPreview";

function Teacher({ teacher }) {
  return (
    <Card>
      <CardContent>
        <UserPreview user={teacher.user} />
      </CardContent>
    </Card>
  );
}

export default Teacher;
