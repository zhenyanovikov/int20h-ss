import { Card, CardContent } from "@mui/material";
import UserPreview from "../UserPreview/UserPreview";

function Student({ student }) {
  return (
    <Card>
      <CardContent>
        <UserPreview user={student.user} />
      </CardContent>
    </Card>
  );
}

export default Student;
