import { gql } from 'apollo-server';
import { GraphQLScalarType } from 'graphql';
import axios from 'axios';

export const typeDefs = gql`
  scalar DateTime

  type Note {
    title: String
    description: String
    date: DateTime
  }

  type Query {
    notes: [Note]
  }
`;

export const resolvers = {
  Query: {
    notes: async () => {
      const res = await axios.get("http://localhost:8080/notes");
      return res.data;
    }
  },
  DateTime: new GraphQLScalarType ({
    name: 'DateTime',
    description: 'Date custom scalar type',
    serialize(value) {
      return value; // Convert outgoing Date to integer for JSON
    },
    parseValue(value) {
      return new Date(value); // Convert incoming integer to Date
    },
    parseLiteral(ast) {
      if (ast.kind === Kind.INT) {
        return new Date(parseInt(ast.value, 10)); // Convert hard-coded AST string to integer and then to Date
      }
      else {
        return new Date(ast.value);
      }
    },
  })
};
