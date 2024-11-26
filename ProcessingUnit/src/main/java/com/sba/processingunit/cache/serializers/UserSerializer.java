package com.sba.processingunit.cache.serializers;

import com.hazelcast.nio.ObjectDataInput;
import com.hazelcast.nio.ObjectDataOutput;
import com.hazelcast.nio.serialization.StreamSerializer;
import com.sba.processingunit.data.User;

import java.io.IOException;

public class UserSerializer implements StreamSerializer<User> {
    @Override
    public void write(ObjectDataOutput output, User user) throws IOException {
        output.writeString(user.getId());
        output.writeString(user.getName());
        output.writeDouble(user.getBalance());
    }

    @Override
    public User read(ObjectDataInput input) throws IOException {
        return new User(input.readString(), input.readString(), input.readDouble());
    }

    @Override
    public int getTypeId() {
        return 1;
    }
}