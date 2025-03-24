import Button from "../../components/ui/Button";

const Landing = () => {
  return (
    <div className="p-8 text-center text-xl font-bold">
      Welcome to <span className="text-blue-500">Landing</span>
      <div className="space-y-4 mt-6">
        <Button
          onClick={() => console.log("Create")}
          variant="primary"
          size="lg"
        >
          Create Room
        </Button>

        <Button variant="danger" size="sm" disabled>
          Abort Mission
        </Button>

        <Button variant="terminal">Join</Button>
      </div>
    </div>
  );
};

export default Landing;
