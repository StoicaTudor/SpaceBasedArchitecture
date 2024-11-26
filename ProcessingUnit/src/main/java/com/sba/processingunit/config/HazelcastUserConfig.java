package com.sba.processingunit.config;

import com.hazelcast.config.Config;
import com.hazelcast.config.MapConfig;
import com.hazelcast.core.Hazelcast;
import com.hazelcast.core.HazelcastInstance;
import org.springframework.stereotype.Component;

@Component
public class HazelcastUserConfig {
    public static final String USERS = "users";
    private final HazelcastInstance hazelcastInstance = Hazelcast.newHazelcastInstance(createConfig());

    public HazelcastInstance getInstance() {
        return hazelcastInstance;
    }
//    public User put(User user) {
//        IMap<String, User> map = hazelcastInstance.getMap(USERS);
//        return map.putIfAbsent(user.getId(), user);
//    }
//
//    public User get(String key) {
//        IMap<String, User> map = hazelcastInstance.getMap(USERS);
//        return map.get(key);
//    }

    private Config createConfig() {
        Config config = new Config().addMapConfig(mapConfig()).setClusterName("dev");
        config.getManagementCenterConfig().setDataAccessEnabled(true).setConsoleEnabled(true).setScriptingEnabled(true);
        config.getNetworkConfig().setPublicAddress("127.0.0.1");
        return config;
        //.getSerializationConfig().addSerializerConfig(serializerConfig())
    }

//    private SerializerConfig serializerConfig() {
//        return new SerializerConfig().setImplementation(new UserSerializer()).setTypeClass(User.class);
//    }

    private MapConfig mapConfig() {
        // Set TTL and Max Idle to -1, which disables expiration
        // No TTL, data never expires
        // No max idle time, data never expires due to inactivity
        return new MapConfig(USERS).setTimeToLiveSeconds(3600).setMaxIdleSeconds(3600);
    }
}